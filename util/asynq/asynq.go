package asynq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/hibiken/asynq"
)

type Server struct {
	Client    *asynq.Client
	Mux       *asynq.ServeMux
	Scheduler *asynq.Scheduler
}

func (s *Server) CreatePushTask(ctx context.Context) {
	task, _ := NewEmailDeliveryTask(42, "some:template:id")
	s.Client.Enqueue(task, asynq.ProcessIn(1*time.Minute), asynq.Unique(time.Minute*3))
	log.Println(ctx, "enqueue task: %s successfully", task.Type())
	s.Mux.HandleFunc(TypeEmailDelivery, HandleEmailDeliveryTask)
}

func (s *Server) CreatePeriodPushTask(ctx context.Context) {
	task, _ := NewStartCron("perform_at", time.Second*30)
	s.Client.Enqueue(task, asynq.ProcessIn(30*time.Second))
	log.Println(ctx, "enqueue task: %s successfully", task.Type())
	stub := &Stub{
		Mux:       s.Mux,
		Client:    s.Client,
		Scheduler: s.Scheduler,
	}
	s.Mux.HandleFunc(TypeStartCron, stub.HandleStartCron)
}

func asyncq_run() {
	s := &Server{}
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	s.Mux = mux
	// ...register other handlers...

	go func() {
		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	}()

	scheduler := asynq.NewScheduler(asynq.RedisClientOpt{Addr: redisAddr}, nil)
	s.Scheduler = scheduler
	go func() {
		if err := scheduler.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	s.Client = client

	s.CreatePushTask(context.Background())
	s.CreatePushTask(context.Background())
	s.CreatePushTask(context.Background())
	// s.CreatePeriodPushTask(context.Background())
	time.Sleep(time.Minute * 10)
}

/*
a simple task with no logic
*/

const typeSimpleTask = "simple_task"

type simpleTaskPayload struct {
	Msg string
}

func newSimpleTask() *asynq.Task {
	bt, _ := json.Marshal(&simpleTaskPayload{Msg: "a simple task"})
	return asynq.NewTask(typeSimpleTask, bt)
}

/*
a simple handler
*/

func simpleHandler(ctx context.Context, t *asynq.Task) error {
	spew.Dump("start to handle...")
	payload := &simpleTaskPayload{}
	if err := json.Unmarshal(t.Payload(), payload); err != nil {
		return err
	}

	spew.Dump(payload.Msg)

	return nil
}

/*
enque task to redis
run
*/

const redisAddr = "127.0.0.1:6379"

var (
	client *asynq.Client
	mux    *asynq.ServeMux
)

func bindHandlerAndTask(delay time.Duration) {
	spew.Dump("delay ", delay)
	spew.Dump("first enqueue")
	taskInfo, err := client.Enqueue(newSimpleTask(), asynq.ProcessIn(delay), asynq.Unique(time.Until(time.Now().Add(30*time.Minute))))
	if err != nil {
		fmt.Println(taskInfo, err)
	}

	spew.Dump("second enqueue")
	taskInfo, err = client.Enqueue(newSimpleTask(), asynq.ProcessIn(delay), asynq.Unique(time.Until(time.Now().Add(30*time.Minute))))
	if err != nil {
		fmt.Println(taskInfo, err)
	}

	spew.Dump("third enqueue")
	taskInfo, err = client.Enqueue(newSimpleTask(), asynq.ProcessIn(delay), asynq.Unique(time.Until(time.Now().Add(30*time.Minute))))
	if err != nil {
		fmt.Println(taskInfo, err)
	}

	spew.Dump("info: enqueue success")
}

func RunAsynq() {
	client = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{},
	)
	mux = asynq.NewServeMux()
	mux.HandleFunc(typeSimpleTask, simpleHandler)
	spew.Dump("info: handler success")

	beginAt := time.Date(2021, 11, 04, 19, 8, 0, 0, time.Local)
	bindHandlerAndTask(time.Until(beginAt))

	go func() {
		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	}()

	time.Sleep(10 * time.Minute)
}

package asynq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	// 定时任务
	TypeEmailDelivery   = "email:deliver"
	// 定时开启周期任务
	TypeStartCron     = "start_cron"
	// 周期任务
	TypeCronTask = "cron_task"
)

type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
}

type StartCrontPayload struct {
	Name string
	Period time.Duration
}

type CronTaskPayload struct {

}

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewEmailDeliveryTask(userID int, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{UserID: userID, TemplateID: tmplID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func NewStartCron(name string, period time.Duration) (*asynq.Task, error) {
	payload, err := json.Marshal(StartCrontPayload{Name: name, Period: period})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeStartCron, payload), nil
}

func NewCronTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CronTaskPayload{})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCronTask, payload), nil
}

//---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.
// Note that it satisfies the asynq.HandlerFunc interface.
//
// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface. See examples below.
//---------------------------------------------------------------

func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	log.Printf("start handling task: %s", t.Type())
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Email to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID)
	// Email delivery code ...
	return nil
}

type Stub struct {
	Client *asynq.Client
	Mux *asynq.ServeMux
	Scheduler *asynq.Scheduler
}


func (s *Stub) HandleStartCron(ctx context.Context, t *asynq.Task) error {
	log.Printf("start handling task: %s", t.Type())
	var p StartCrontPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	task,_ := NewCronTask()
	s.Scheduler.Register("@every 30s", task)
	s.Mux.HandleFunc(TypeCronTask, HandleCronTask)

	return nil
}

func HandleCronTask(ctx context.Context, t *asynq.Task) error {
	log.Printf("start handling period task: %s", t.Type())
	return nil
}
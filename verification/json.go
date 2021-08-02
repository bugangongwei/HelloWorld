package verification

import (
	"encoding/json"
	"fmt"
)

type JsonTag struct {
	Name string `json:"name_str"`
	Age  int    `json:"age_int"`
}

func UnmarshalJson() {
	jt := &JsonTag{}

	if err := json.Unmarshal([]byte(`{"name_str":"susu","age_int":16}`), jt); err != nil {
		panic(err)
	}
	fmt.Println(jt)

	jtt := &JsonTag{}
	if err := json.Unmarshal([]byte(`{"nameStr":"susu","ageInt":16}`), jtt); err != nil {
		panic(err)
	}
	fmt.Println(jtt)
}

type Data struct {
	AppId string          `json:"app_id"`
	Extra json.RawMessage `json:"extra"`
}

func UnmarshalToRaw() {
	text := `{"app_id":"lls","extra":{"audio_url":"http://cdn.llsapp.com/5eb4c383636f6e000a0002aa_fdc7e6699a40af01_1625099162094.mp3","activity_id":"5eb4c383636f6e000a0002ab","text":"潮落之后必有潮起，下次说得更好肯定没问题。来听听我的口语吧。","score":81,"scores":[{"word_scores":[78,80,94],"sentence_score":89,"duration":2080},{"word_scores":[85,88,89,76,78,72,88,49,77],"sentence_score":84,"duration":3360},{"word_scores":[55],"sentence_score":55,"duration":1520},{"word_scores":[76,76,82],"sentence_score":85,"duration":1920},{"word_scores":[79,85,64,65,54,55,68,73,92],"sentence_score":77,"duration":5040},{"word_scores":[52,89,11,78,93],"sentence_score":82,"duration":2640},{"word_scores":[66,96,81],"sentence_score":85,"duration":1920},{"word_scores":[90,84,91,77,86,89,96],"sentence_score":91,"duration":3440},{"word_scores":[79,83,97,72,61],"sentence_score":86,"duration":2640}],"user":{"id":"OGZjMWQwMDAwYjQxY2EzZg==","nick":"张铭扬","avatar":"https://cdn.llscdn.com/fdc7e6699a40af01_1617323487651.jpg"},"created_at":1625098928,"course_name":"“背锅甩锅”英语怎么说：“背锅”英语怎么说","share_image_url":"","recommendation":""}}`

	data := &Data{}
	if err := json.Unmarshal([]byte(text), data); err != nil {
		panic(err)
	}
	fmt.Println(string(data.Extra))

}

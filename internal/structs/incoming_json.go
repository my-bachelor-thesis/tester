package structs

type IncomingJson struct {
	Solution string `json:"solution"`
	Test     string `json:"test"`
}

//func (s IncomingJson) Check() error {
//	if s.Code == "" {
//		return errors.New("code field cannot be empty")
//	}
//	if s.TaskId < 0 {
//		return errors.New("task_id field has to be a positive integer")
//	}
//
//	if s.Type == "code" {
//		if s.UserId < 0 {
//			return errors.New("user_id field has be a positive integer")
//		}
//	} else if s.Type != "test" {
//		return errors.New("type field has to code or test")
//	}
//	return nil
//}

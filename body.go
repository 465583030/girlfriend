package gf

type Array []interface{}
type Object map[string]interface{}

func (req *Request) RemoveKey(key string) { delete(req.Object, key) }

func (req *Request) Value(key string) (bool, interface{}) {
	for {
		if req.Object[key] == nil { req.Debug("VALUE FOR THIS KEY IS NIL"); break }
		return true, req.Object[key]
	}
	return false, nil
}

// map[string]interface{}
func (req *Request) MSI(key string) (bool, Object) {
	for {
		if req.Object[key] == nil { req.Debug("VALUE FOR THIS KEY IS NIL"); break }
		s, ok := req.Object[key].(map[string]interface{}); if !ok { req.Reflect(req.Object[key]); break }
		return true, s
	}
	return false, nil
}

// interface array
func (req *Request) IA(key string) (bool, []interface{}) {
	for {
		if req.Object[key] == nil { req.Debug("VALUE FOR THIS KEY IS NIL"); break }
		s, ok := req.Object[key].([]interface{}); if !ok { req.Reflect(req.Object[key]); break }
		return true, s
	}
	return false, nil
}

func (req *Request) String(key string) (bool, string) {
	for {
		if req.Object[key] == nil { req.Debug("VALUE FOR THIS KEY IS NIL"); break }
		s, ok := req.Object[key].(string); if !ok { req.Reflect(req.Object[key]); break }
		return true, req.BlueMonday.Sanitize(s)
	}
	return false, ""
}

func (req *Request) Float64(key string) (bool, float64) {
	for {
		if req.Object[key] == nil { req.Debug("VALUE FOR THIS KEY IS NIL"); break }
		s, ok := req.Object[key].(float64); if !ok { req.Reflect(req.Object[key]); break }
		return true, s
	}
	return false, 0.0
}

func (req *Request) Int(key string) (bool, int) {
	for {
		if req.Object[key] == nil { req.Debug("VALUE FOR THIS KEY IS NIL"); break }
		s, ok := req.Object[key].(float64); if !ok { req.Reflect(req.Object[key]); break }
		return true, int(s)
	}
	return false, 0
}

func (req *Request) Int64(key string) (bool, int64) {
	for {
		if req.Object[key] == nil { req.Debug("VALUE FOR THIS KEY IS NIL"); break }
		s, ok := req.Object[key].(float64); if !ok { req.Reflect(req.Object[key]); break }
		return true, int64(s)
	}
	return false, 0
}

func (req *Request) Bool(key string) (bool, bool) {
	for {
		if req.Object[key] == nil { req.Debug("VALUE FOR THIS KEY IS NIL"); break }
		s, ok := req.Object[key].(bool); if !ok { req.Reflect(req.Object[key]); break }
		return true, s
	}
	return false, false
}

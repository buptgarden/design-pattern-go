package proxy

type Application struct{}

func (a *Application) Handlerequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "ok"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}

	return 404, "Not Ok"
}

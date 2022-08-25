package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/
var EventId = 1

// Структура пользователя
type User struct {
	ID int
}

//Структура события
type Event struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	User        User      `json:"user"`
}

//Структура результата
type Result struct {
	Result []Event `json:"result"`
}

//Структура ошибки
type Errors struct {
	Error []string `json:"error"`
}

var Events = make(map[int][]Event)

// Добавляем событие
func AddEvent(ev Event) Event {
	ev.Id = EventId
	Events[ev.Id] = append(Events[ev.Id], ev)
	EventId++
	return ev
}

// Обновляем событие
func UpdateEvent(ev Event) Errors {
	var er Errors
	if _, inMap := Events[ev.Id]; inMap {
		Events[ev.Id] = append(Events[ev.Id], ev)
	} else {
		e := "Не возможно обновить событие, так как его не существует"
		er.Error = append(er.Error, e)
	}
	return er
}

//Удаляем событие
func DeleteEvent(ev Event) Errors {
	var er Errors
	if _, inMap := Events[ev.Id]; inMap {
		delete(Events, ev.Id)
	} else {
		e := "Не возможно удалить событие, так как его не существует"
		er.Error = append(er.Error, e)
	}

	return er
}

//Создаем событие
func CreateEvent(r *http.Request) (Event, Errors) {
	var e Event
	var er Errors
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	desc := r.FormValue("description")
	if len(desc) <= 255 && len(desc) >= 3 {
		e.Description = desc
	} else {
		e := "Описание должно быть более 3 символов и менее 255"
		er.Error = append(er.Error, e)

	}

	t, _ := time.Parse("2006-01-02", r.FormValue("date"))
	e.Date = t
	var u User
	u.ID, _ = strconv.Atoi(r.FormValue("user_id"))
	if u.ID > 0 {
		e.User = u
	} else {
		e := "Id пользователя не может быть меньше 0"
		er.Error = append(er.Error, e)
	}

	return e, er
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	e, err := CreateEvent(r)
	if len(err.Error) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(w.Write(StructToJson(err)))

	} else {
		e = AddEvent(e)
		log.Println(w.Write(StructToJson(e)))
	}
}

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	e, er := CreateEvent(r)

	if len(er.Error) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(w.Write(StructToJson(er)))

	} else {
		err := UpdateEvent(e)
		if len(err.Error) > 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Println(w.Write(StructToJson(err)))
		}
	}
}
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	e, er := CreateEvent(r)
	if len(er.Error) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(w.Write(StructToJson(er)))

	} else {
		err := DeleteEvent(e)
		if len(err.Error) > 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Println(w.Write(StructToJson(err)))

		}
	}
}

func GetEventsForWeek() Result {
	var result Result
	var res []Event
	t := time.Now()
	past := t.AddDate(0, 0, -3)
	future := t.AddDate(0, 0, 4)
	for _, r := range Events {
		for _, k := range r {
			if k.Date.Before(future) && k.Date.After(past) {
				res = append(res, k)
			}
		}
	}

	result.Result = res
	return result

}

func GetEventsForMonth() Result {
	var result Result
	var res []Event
	t := time.Now()
	past := t.AddDate(0, -15, 0)
	future := t.AddDate(0, 15, 0)
	for _, r := range Events {
		for _, k := range r {
			if k.Date.Before(future) && k.Date.After(past) {
				res = append(res, k)
			}
		}
	}
	result.Result = res
	return result
}

func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	res := GetEventsForWeek()
	re := StructToJson(res)
	log.Println(w.Write(re))

}
func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	res := GetEventsForMonth()
	re := StructToJson(res)
	log.Println(w.Write(re))
}

func JsonResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})

}
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(time.Since(start))
	})

}

func main() {

	// Обработчики...
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", CreateEventHandler)
	mux.HandleFunc("/update_event", UpdateEventHandler)
	mux.HandleFunc("/delete_event", DeleteEventHandler)
	mux.HandleFunc("/events_for_week", EventsForWeek)
	mux.HandleFunc("/events_for_month", EventsForMonth)

	handler := LoggerMiddleware(JsonResponseMiddleware(mux))

	//Сервер...
	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server is listening...")
	log.Fatal(s.ListenAndServe())
}

// Перевод любой структуры в Json
func StructToJson(str any) []byte {
	Json, err := json.Marshal(str)
	if err != nil {
		fmt.Println(err)
	}
	return Json
}

// Перевод любого джейсона в структуру
func JsonToStruct[T any](jsn []byte, strct T) T {
	if err := json.Unmarshal(jsn, &strct); err != nil {
		fmt.Println(err)
	}
	return strct
}

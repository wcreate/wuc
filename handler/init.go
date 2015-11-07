package handler

import "gopkg.in/macaron.v1"

func InitHandles(m *macaron.Macaron) {
  m.Delete("/api/user/:uid", DeleteUser)

  m.Get("/api/user/info/:uid", UserInfo)
  m.Put("/api/user/info/:uid", ModifyUser)
  m.Post("/api/user/info/:uid", ModifyUser)
  m.Put("/api/user/pwd/:uid", ModifyPassword)

  m.Post("/api/user/login", LoginUser)
}

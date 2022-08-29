### Microservice demo project using Python, Golang and PHP
<img src="https://github.com/phammyan530/golang-api-server/blob/main/image/microservice-demo-project.jpg">

#### This app is a demo of using: 

- Python scraping website into MySQL server: https://github.com/phammyan530/python-scraping-website-into-mysql
- PHP powered frontend, a Top 250 Movies website: https://github.com/phammyan530/codeigniter-3-movie-blog
- Golang powered backend: REST API application:

### REST based microservices API development in Golang
This app is a demo of using a Golang powered backend:
- [Echo framework](https://echo.labstack.com/): High performance, extensible, minimalist Go web framework.
- [Beego ORM](https://beego.vip/docs/mvc/model/overview.md): Beego ORM is a powerful ORM framework written in Go.
- Group API
- JWT authentication

### API list

| Method | URI                     | Note                                  |
|--------|-------------------------|---------------------------------------|
| GET    | /api/movie/get_all{page}| get list movie with pagination        |
| GET    | /api/movie/get{id}      | get detail movie by Movie ID          |
| GET    | /api/movie/get_random   | get list random movie                 |
| GET    | /api/admin/list_movie   | authentication to admin page          |
| POST   | /api/admin/update_movie | authentication to update detail movie |
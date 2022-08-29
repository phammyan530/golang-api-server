package handler

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type Movies struct {
	Id          int64  `orm:"auto" json:"id"`
	Movie       string `json:"movie"`
	Year        string `json:"year"`
	Run_time    string `json:"run_time"`
	Genre       string `json:"genre"`
	Director    string `json:"director"`
	Actors      string `json:"actors"`
	Descrt      string `json:"descrt"`
	Img         string `json:"img"`
	Total_gross string `json:"total_gross"`
	Rating      string `json:"rating"`
	Total_rate  string `json:"total_rate"`
}

func init() {
	orm.RegisterModel(new(Movies))
}

func GetAllMovies(c echo.Context) error {
	page := cast.ToInt(c.QueryParam("page"))
	offset := page * 10

	var movies []Movies
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("movies.*",
		"reviews.rating", "reviews.total_rate",
		"gross.total_gross").
		From("movies").
		InnerJoin("gross").On("movies.id = gross.movie_id").
		InnerJoin("reviews").On("movies.id = reviews.movie_id").
		OrderBy("id").Asc().
		Limit(10).Offset(offset)
	sql := qb.String()
	o := orm.NewOrm()
	num, err := o.Raw(sql).QueryRows(&movies)

	if err != nil {
		glog.Errorf("\nError get all movies %v: %v", num, err)
		return err
	}
	return c.JSON(http.StatusOK, movies)
}

func GetMovie(c echo.Context) error {
	id := cast.ToInt64(c.QueryParam("id"))
	if id == 0 {
		id = 1
	}

	var movie Movies
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("movies.*",
		"reviews.rating", "reviews.total_rate",
		"gross.total_gross").
		From("movies").
		InnerJoin("gross").On("movies.id = gross.movie_id").
		InnerJoin("reviews").On("movies.id = reviews.movie_id").
		Where("movies.id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	err := o.Raw(sql, id).QueryRow(&movie)

	if err != nil {
		glog.Errorf("\nError get 1 movies: %v", err)
		return err
	}
	return c.JSON(http.StatusOK, movie)
}

func GetRandomMovies(c echo.Context) error {
	rand.Seed(time.Now().UnixNano())
	var idlist = randIntList(1, 250, 15)

	var movies []Movies
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("movies.*",
		"reviews.rating", "reviews.total_rate",
		"gross.total_gross").
		From("movies").
		InnerJoin("gross").On("movies.id = gross.movie_id").
		InnerJoin("reviews").On("movies.id = reviews.movie_id").
		Where("movies.id").In(idlist)
	sql := qb.String()
	o := orm.NewOrm()
	num, err := o.Raw(sql).QueryRows(&movies)

	if err != nil {
		glog.Errorf("\nError get random movies %v: %v", num, err)
		return err
	}
	return c.JSON(http.StatusOK, movies)
}

func randIntList(min, max int, n int) string {
	res := make([]string, n)
	for i := range res {
		res[i] = strconv.Itoa(min + rand.Intn(max-min))
	}
	return strings.Join(res, ",")
}

func UpdateMovie(c echo.Context) error {
	id := cast.ToInt(c.QueryParam("id"))
	if id <= 0 {
		return nil
	}

	movie_name := c.QueryParam("movie_name")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("movies").
		Set("movie = ?").
		Where("id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	_, err := o.Raw(sql, movie_name, id).Exec()

	if err != nil {
		glog.Errorf("\nError on update movies: %v", err)
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

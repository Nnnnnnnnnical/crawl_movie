package controllers

import (
	"crawl_movie/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"strconv"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie(){
	var movieInfo models.MovieInfo

	//连接url
	models.ConnectRedis("127.0.0.1:6379")

	//爬虫的入口url
	sUrl := "https://movie.douban.com/subject/1302434/?from=subject-page"
	models.PutinQueue(sUrl)


	for{
		length := models.GetQueueLength()
		if length == 0{
			break
		}

		sUrl = models.PopfromQueue()

		//判断sUrl是否被访问过了
		if models.IsVisit(sUrl){
			continue
		}

		rsp := httplib.Get(sUrl)
		sMovieHtml,err := rsp.String()
		if err!=nil{
			panic(err)
		}
		movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
		//记录电影信息
		//if movieInfo.Movie_name != ""{
			//c.Ctx.WriteString("<br>"+movieInfo.Movie_name+"</br>")
			movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_writer = models.GetMovieWriter(sMovieHtml)
			movieInfo.Movie_main_character = models.GetMovieMainCharacters(sMovieHtml)
			movieInfo.Movie_type = models.GetMovieType(sMovieHtml)
			movieInfo.Movie_country = models.GetMovieCountry(sMovieHtml)
			movieInfo.Movie_language = models.GetMoiveLanguage(sMovieHtml)
			movieInfo.Movie_on_time = models.GetMovieOnTime(sMovieHtml)
			movieInfo.Movie_span = models.GetMovieSpan(sMovieHtml)
			movieInfo.Movie_grade = models.GetMovieGrade(sMovieHtml)
			movieInfo.Movie_id, _ = strconv.ParseInt(models.GetMovieId(sMovieHtml),10,64)
			movieInfo.Movie_pic = models.GetMoviePic(sMovieHtml)

			models.AddMovie(&movieInfo)

		//}
		//提取该页面的所有连接
		urls := models.GetMovieUrls(sMovieHtml)

		for _,url := range urls{
			models.PutinQueue(url)
		}

		//sUrl应该记录到set中
		models.AddToSet(sUrl)


	}

	c.Ctx.WriteString("end of crawl")

}

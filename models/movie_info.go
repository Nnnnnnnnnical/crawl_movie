package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"regexp"
)

var (
	db orm.Ormer
)

type MovieInfo struct {
	Id int64
	Movie_id int64
	Movie_name string
	Movie_pic string
	Movie_director string
	Movie_writer string
	Movie_country string
	Movie_language string
	Movie_main_character string
	Movie_type string
	Movie_on_time string
	Movie_span string
	Movie_grade string
	_Create_time string
}

func init(){
	orm.Debug = true;
	orm.RegisterDataBase("default","mysql","root:@tcp(127.0.0.1:3306)/crawl_movie")
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}

func AddMovie(movie_info *MovieInfo)(int64,error)  {
	movie_info.Id = 0
	id,err :=db.Insert(movie_info);
	return id,err;

}

func GetMovieDirector(movieHtml string) string{
	if movieHtml == ""{
		return ""
	}

	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result) == 0{
		return ""
	}
	return string(result[0][1])
}

func GetMovieName(movieHtml string)string  {
	if movieHtml == ""{
		return ""
	}

	reg := regexp.MustCompile(`<span property="v:itemreviewed">(.*)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])

}
func GetMovieWriter(movieHtml string)string  {
	if movieHtml == ""{
		return ""
	}
	reg := regexp.MustCompile(`<a.*?href="/celebrity/.*?/">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result) == 0{
		return ""
	}
	writer :=  ""
	for _,v :=range result{
		writer +=v[1] + "/"
	}
	return writer

}


func GetMovieMainCharacters(movieHtml string)string  {
	if movieHtml == ""{
		return ""
	}

	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)

	mainCharacters :=  ""
	for _,v :=range result{
		mainCharacters +=v[1] + "/"
	}
	return mainCharacters

}

func GetMovieType(movieHtml string)string  {
	if movieHtml == ""{
		return ""
	}

	reg := regexp.MustCompile(`<span property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result) == 0{
		return ""
	}
	movieType :=  ""
	for _,v :=range result{
		movieType +=v[1] + "/"
	}
	return movieType

}

func GetMovieCountry(movieHtml string)string  {
	if movieHtml == ""{
		return ""
	}

	reg := regexp.MustCompile(`<span class="pl">制片国家/地区:</span>(.*)<br/>`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result) == 0{
		return ""
	}
	return string(result[0][1])

}


func GetMoiveLanguage(movieHtml string)string  {
	if movieHtml == ""{
		return ""
	}

	reg := regexp.MustCompile(`<span class="pl">语言:</span> (.*)<br/>`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result) == 0{
		return ""
	}
	return string(result[0][1])

}

func GetMovieOnTime(movieHtml string)string  {
	if movieHtml == ""{
		return ""
	}


	reg := regexp.MustCompile(`<span class="pl">上映日期:</span> <span property="v:initialReleaseDate" .*?>(.*)</span><br/>`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result) == 0{
		return ""
	}
	return string(result[0][1])

}

func GetMovieSpan(movieHtml string)string  {
	if movieHtml == ""{
		return ""
	}


	reg := regexp.MustCompile(`<span class="pl">片长:</span> <span property="v:runtime".*?>(.*)</span><br/>`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result) == 0{
		return ""
	}
	return string(result[0][1])

}

func GetMovieGrade(movieHtml string)string  {
	if movieHtml == ""{
		return ""
	}


	reg := regexp.MustCompile(`<strong class="ll rating_num" property="v:average">(.*)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result) == 0{
		return ""
	}
	return string(result[0][1])

}

func GetMovieId(movieHtml string)string  {
	if movieHtml == ""{
		return ""
	}


	reg := regexp.MustCompile(`<span class="rec" id="电影-(.*)">`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result) == 0{
		return ""
	}
	return string(result[0][1])

}

func GetMoviePic(movieHtml string)string  {
	if movieHtml == ""{
		return ""
	}


	reg := regexp.MustCompile(`<img src="(.*)" title="点击看更多海报" alt=.*? rel="v:image" />`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)
	if len(result) == 0{
		return ""
	}
	return string(result[0][1])

}

func GetMovieUrls(movieHtml string)[]string  {

	reg := regexp.MustCompile(`<a href="(https://movie.douban.com/.*?)" class="" >`)
	result := reg.FindAllStringSubmatch(movieHtml,-1)

	var movieSets []string
	for _,v := range result{
		movieSets = append(movieSets,v[1])
	}

	return movieSets

}


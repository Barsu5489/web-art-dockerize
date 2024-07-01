package asciiweb

import (
	"html/template"
	"net/http"

	art "asciiweb/server/ascii-art"
	// "fmt"
)

type Result struct {
	Data string
}

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	var artOutput Result
	var res string
	pageData := `
	<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <script src="https://cdn.tailwindcss.com"></script>
  <title>Ascii Art Web</title>
</head>

<body>
  <div class="container relative mx-auto p-6">

    <h1 class="text-slate-800 text-2xl">Ascii Art Web</h1>
    <div class="flex flex-row my-6 ">
      <form action="/" method="post">
        <textarea id="message" rows="4" cols="80"
          class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="Enter your text here..." name="text" id="text"></textarea>

        <div class="flex flex-col my-6 ">
            <p class="text-slate-600">Choose banner file you want to use.</p>
          <div class="items-center text-gray-600 py-2 ">
            <input type="radio" name="banner" id="standard" value="standard.txt" checked>
            <label for="standard">standard</label>
          </div>
          <div class="items-center text-gray-600 py-2" >
            <input type="radio" name="banner" id="thinkertoy" value="thinkertoy.txt">
            <label for="thinkertoy">thinkertoy</label>
          </div>
          <div class="items-center text-gray-600 py-2">
            <input type="radio" name="banner" id="shadow" value="shadow.txt">
            <label for="shadow">shadow</label>
          </div>
          <input type="submit" class="h-auto w-[200px] border py-2 px-4 text-center rounded-lg bg-blue-500  text-white"
            value="Create art">
        </div>
      </form>
    </div>

	<fieldset>
	<legend>Generated Art</legend>
    <pre class="py-6">{{.Data}}</pre>
	</fieldset>
    <footer class="bg-white rounded-lg  m-4 sticky bottom-0">
       <span class="block text-sm text-gray-500 sm:text-center dark:text-gray-400">© 2024 <a href="/" class="hover:underline">Ascii Art Web™</a>. All Rights Reserved.</span>
      </footer>
  </div>
</body>

</html>
`
	r.ParseForm()
	userInput := r.FormValue("text")
	bannerFile := r.FormValue("banner")

	if userInput == "" {
		artOutput = Result{
			Data: "User Input Cannot Be Empty.",
		}
		tpl, _ := template.Must(template.ParseFiles("templates/index.html")).Parse(pageData)

		tpl.Execute(w, artOutput)
		return
	}

	res = art.GenArt(userInput, bannerFile)
	if res == "" {
		artOutput = Result{
			Data: "Bad Request",
		}

	} else {
		artOutput = Result{
			Data: res,
		}
	}

	tpl, _ := template.Must(template.ParseFiles("templates/index.html")).Parse(pageData)

	tpl.Execute(w, artOutput)
}

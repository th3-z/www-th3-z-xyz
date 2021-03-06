function editArticle(name) {
    console.log(name)

    emContent = document.getElementById("article_content")
    // TODO: show/hide em instead
    emContent.innerHTML = "<textarea id=\"article_edit\" style=\"width: 100%; height: 80em;\"> Loading... </textarea>"

    getSrc(name)
}

function saveArticle(name) {
    console.log(name)
}

function deleteArticle(name) {
    console.log(name)
}

function getSrc(name) {
    $.ajax({
        method: "POST",
        url: "/articles/" + name + "/src",
        success: function(data) {

            emEdit = document.getElementById("article_edit")
            emEdit.innerHTML = data
        }
    });
}
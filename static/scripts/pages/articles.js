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
    $.ajax({
        method: "POST",
        url: "/articles/" + name + "/delete",
        success: function(data) {
            alert("article deleted")
        }
    });
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

function visibleArticle(name) {
    let btnVisible = document.getElementById("btn_visible")
    let isVisible = btnVisible.dataset.visible != "0"

    $.ajax({
        method: "POST",
        url: "/articles/" + name + "/set_visible",
        data: {
            visible: !isVisible
        },
        success: function() {
            if (isVisible) {
                btnVisible.innerHTML = "<i class=\"fas fa-eye-slash\">"
                btnVisible.dataset.visible = "0"
            } else {
                btnVisible.innerHTML = "<i class=\"fas fa-eye\">"
                btnVisible.dataset.visible = "1"
            }
        }
    });
}
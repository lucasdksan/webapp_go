$("#new-post").on("submit", create_post);

function create_post(e){
    e.preventDefault();

    const title = $("#title").val();
    const content = $("#content").val();

    $.ajax({
        url: "/publications",
        method: "POST",
        data: {
            title,
            content,
        }
    }).done(()=>{
        window.location = "/home";
    }).fail((err)=>{
        console.log("Error: ", err);
        alert("Erro ao cadastrar o post!");
    });
}
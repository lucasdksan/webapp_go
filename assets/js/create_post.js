$("#new-post").on("submit", create_post);
$(".like-post-btn").on("click", like_post);

function like_post(event){
    const event_click = $(event.target);
    const publication_id = $(event_click).closest("div.post-card").data("publication-id");

    $.ajax({
        url: `/publications/${publication_id}/like`,
        method: "POST"
    }).done(()=>{
        alert("Curtida")
    }).fail((err)=>{
        console.error("Erro na curtida: ", err)
    });
}

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
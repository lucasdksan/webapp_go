$("#new-post").on("submit", create_post);
$(document).on("click", ".liked-post", like_post);
$(document).on("click", ".disliked-post", dislike_post);

function dislike_post(event){
    const event_click = $(event.target);
    const publication_id = $(event_click).closest("div.post-card").data("publication-id");
    
    event_click.prop("disabled", true);

    $.ajax({
        url: `/publications/${publication_id}/dislike`,
        method: "POST"
    }).done(()=>{
        const count_likes = event_click.next("span");
        const likes_val = parseInt(count_likes.text());

        count_likes.text(likes_val <= 0 ? 0 : likes_val - 1);
        event_click.addClass("liked-post").removeClass("disliked-post").removeClass("text-danger");
    }).fail((err)=>{
        console.error("Erro na curtida: ", err)
    }).always(()=>{
        event_click.prop("disabled", false);
    });
}

function like_post(event){
    const event_click = $(event.target);
    const publication_id = $(event_click).closest("div.post-card").data("publication-id");

    event_click.prop("disabled", true);
    $.ajax({
        url: `/publications/${publication_id}/like`,
        method: "POST"
    }).done(()=>{
        const count_likes = event_click.next("span");
        const likes_val = parseInt(count_likes.text());

        count_likes.text(likes_val + 1);
        event_click.addClass("disliked-post").removeClass("liked-post").addClass("text-danger");
    }).fail((err)=>{
        console.error("Erro na curtida: ", err)
    }).always(()=>{
        event_click.prop("disabled", false);
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
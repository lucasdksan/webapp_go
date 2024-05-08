$(".delete-post").on("click", delete_post);

function delete_post() {
    const _this = this;
    const card = $(_this).closest("div.post-card");
    const publication_id = card.data("publication-id");

    $(_this).prop("disabled", true);

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Esse processo é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((check)=>{
        if(!check.value) return;

        $.ajax({
            url: `/publications/${publication_id}`,
            method: "DELETE",
        }).done(function(){
            card.fadeOut("slow", function(){
                $(this).remove();
            });
        }).fail(function(e){
            console.error(e);
            Swal.fire({
                title: "Ooops...",
                text: "Erro ao excluir!",
                icon: "error"
            });
        });
    })
}
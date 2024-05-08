$("#update-btn").on("click", update_post);

function update_post(){ 
    $(this).prop("disabled", true);
    
    const _this = this;
    const publication_id = $(this).data("publication-id");
    
    $.ajax({
        url: `/publications/${publication_id}`,
        method: "PUT",
        data: {
            title: $("#title").val(),
            content: $("#content").val()
        }
    }).done(()=>{
        Swal.fire({
            title: "Sucesso",
            text: "Publicação atualizada com sucesso!",
            icon: "sucess"
        }).then(()=>{
            window.location = "/home";
        });
    }).fail(()=>{
        Swal.fire({
            title: "Ooops...",
            text: "Erro na atualização!",
            icon: "error"
        }).then(()=>{
            window.location = "/home";
        });
    }).always(()=>{
        $(_this).prop("disabled", false);
    });
}
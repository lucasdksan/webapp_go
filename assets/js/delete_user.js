function user_delete(event){
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja apagar a sua conta? Esse é uma irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((check)=>{
        if(check.value){
            $.ajax({
                url: "/user-delete",
                method: "DELETE",
            }).done(()=>{
                Swal.fire({
                    title: "Sucesso!",
                    text: "Sua conta foi deletada com sucesso!",
                    icon: "sucess"
                }).then(()=>{
                    window.location = "/logout";
                });
            }).fail((err)=>{
                console.log("Error: ", err);
                Swal.fire({
                    title: "Ooops...",
                    text: "Erro ao deletar a conta!",
                    icon: "error"
                });
            });
        }
    });
}

$("#user-delete").on("click", user_delete);
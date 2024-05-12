function un_follow(){
    const _this = this;
    const user_id = $(_this).data("user-id");

    $(_this).prop("disabled", true);
    $.ajax({
        url: `/users/${user_id}/unfollow`,
        method: "POST"
    }).done(function(){
        window.location = `/users/${user_id}`; 
    }).fail(function(){
        Swal.fire({
            title: "Ooops...",
            text: "Erro ao parar de seguir o usuário!",
            icon: "error"
        });

        $(_this).prop("disabled", false);
    });
}

function follow(){
    const _this = this;
    const user_id = $(_this).data("user-id");

    $(_this).prop("disabled", true);
    $.ajax({
        url: `/users/${user_id}/follow`,
        method: "POST"
    }).done(function(){
        window.location = `/users/${user_id}`; 
    }).fail(function(){
        Swal.fire({
            title: "Ooops...",
            text: "Erro ao parar de seguir o usuário!",
            icon: "error"
        });

        $(_this).prop("disabled", false);
    });
}

$("#unfollow").on("click", un_follow);
$("#follow").on("click", follow);
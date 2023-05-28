$('#nova-publicacao').on('submit', criarPublicacao)

function criarPublicacao(evento) {
    evento.preventDefault();
    let teste = $('#conteudo').val()
    alert("conteudo:" + teste)
    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function () {
        window.location = "/home";
    }).fail(function () {
        alert("Erro ao criar a publicação!");
    })
}
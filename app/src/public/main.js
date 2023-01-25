document.addEventListener('DOMContentLoaded', () => {
    let loc = window.location;
    let uri = 'ws:';
    if (loc.protocol === 'https:') {
      uri = 'wss:';
    }
    uri += '//' + loc.host;
    uri += loc.pathname + 'ws';
  
    // web socket接続
    const ws = new WebSocket(uri);

    // 接続通知
    ws.onopen = function()  {
      console.log('Connected');
    }

    // エラー発生
    // ws.onerror = function(error) {
    //     document.getElementById( "eventType" ).value = "エラー発生イベント受信";
    //     document.getElementById( "dispMsg" ).value = error.data;
    // };
  
    // メッセージ受信
    ws.onmessage = function(evt) {
      let out = document.getElementById('output');
      out.innerHTML += evt.data + '<br />';
    }
  
    const btn = document.getElementById('btn');
    btn.addEventListener('click', () => {
      ws.send(document.getElementById('input').value);
    });

    // 切断
    // connection.onclose = function() {
    //     document.getElementById( "eventType" ).value = "通信切断イベント受信";
    //     document.getElementById( "dispMsg" ).value = "";
    // };
  });
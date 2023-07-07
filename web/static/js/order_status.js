async function sendRequest(method, url, body) {
    let response = await fetch(url, {
        method: method,
        headers: { 'Content-Type': 'application/json;charset=utf-8' },
        body: JSON.stringify(body)
    });
    data = await response.json();
    return data;
};
document.addEventListener("DOMContentLoaded", function () {

    document.getElementById("orderStatus").style.display = "none";

});
function OrderStatus(el) {
    var idOrder = el.id.value;
    document.getElementById("orderStatus").style.display = "block";
    sendRequest('get', '/api/order/' + idOrder)
        .then(data => {
            document.getElementById("order_uid").innerHTML = data.order_uid;
            document.getElementById("track_number").innerHTML = data.track_number;
            document.getElementById("entry").innerHTML = data.entry;
            document.getElementById("locale").innerHTML = data.locale;
            document.getElementById("internal_signature").innerHTML = data.internal_signature;
            document.getElementById("customer_id").innerHTML = data.customer_id;
            document.getElementById("delivery_service").innerHTML = data.delivery_service;
            document.getElementById("shardkey").innerHTML = data.shardkey;
            document.getElementById("sm_id").innerHTML = data.sm_id;
            document.getElementById("date_created").innerHTML = data.date_created;
            document.getElementById("oof_shard").innerHTML = data.oof_shard;
            document.getElementById("name").innerHTML = data.delivery.name;
            document.getElementById("phone").innerHTML = data.delivery.phone;
            document.getElementById("zip").innerHTML = data.delivery.zip;
            document.getElementById("city").innerHTML = data.delivery.city;
            document.getElementById("address").innerHTML = data.delivery.address;
            document.getElementById("region").innerHTML = data.delivery.region;
            document.getElementById("email").innerHTML = data.delivery.email;
            document.getElementById("transaction").innerHTML = data.payment.transaction;
            document.getElementById("request_id").innerHTML = data.payment.request_id;
            document.getElementById("currency").innerHTML = data.payment.currency;
            document.getElementById("provider").innerHTML = data.payment.provider;
            document.getElementById("amount").innerHTML = data.payment.amount;
            document.getElementById("payment_dt").innerHTML = data.payment.payment_dt;
            document.getElementById("bank").innerHTML = data.payment.bank;
            document.getElementById("delivery_cost").innerHTML = data.payment.delivery_cost;
            document.getElementById("goods_total").innerHTML = data.payment.goods_total;
            document.getElementById("custom_fee").innerHTML = data.payment.custom_fee;
            let html = "";
            if (data !== null) {
                let index;
                for (index = 0; index < data.items.length; ++index) {
                    var item = (data.items[index]);
                    html += (`<tr>
                <td><label>${index + 1}</label></td>
                <td><label>${item.chrt_id}</label></td>
                <td><label>${item.track_number}</label></td>
                <td><label>${item.price}</label></td>
                <td><label>${item.rid}</label></td>
                <td><label>${item.name}</label></td>
                <td><label>${item.sale}</label></td>
                <td><label>${item.size}</label></td>
                <td><label>${item.total_price}</label></td>
                <td><label>${item.nm_id}</label></td>
                <td><label>${item.brand}</label></td>
                <td><label>${item.status}</label></td>
            </tr>  `);
                };
            };
            document.getElementById("items").innerHTML = html;
      
      
        });
   
    return false;
};

var hr = document.getElementById("text")
hr.innerHtml = '<img src='
var jr = document.getElementById("json")
var xhr = new XMLHttpRequest();
xhr.onreadystatechange = function() {
	if (xhr.readyState == XMLHttpRequest.DONE) {
		jr.innerText = xhr.responseText;
		identity=JSON.parse(xhr.responseText);
		var buf = "";
		for (var key in identity) {
			if (identity.hasOwnProperty(key)) {
				buf += '<label class="control-label col-sm-2" >' +
					key +
					'</label><div class="col-sm-10"><p class="form-control-static">' +
					identity[key] +
					'</p></div>'
			}
		}
		hr.innerHtml = buf
	}
}
xhr.open('GET', 'https://identigen.herokuapp.com', true);
xhr.send(null);

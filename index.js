function spacer(key){
	return key.replace(/([A-Z])/g, ' $1').trim()
}
var hr = document.getElementById("campi")
var jr = document.getElementById("json")
function regen(){
var xhr = new XMLHttpRequest();
xhr.onreadystatechange = function() {
	if (xhr.readyState == XMLHttpRequest.DONE) {
		jr.innerText = xhr.responseText;
		identity=JSON.parse(xhr.responseText)[0];
		var buf = "";
		for (var key in identity) {
			if (identity.hasOwnProperty(key)) {
				buf += '<label class="control-label col-sm-5" >' +
					spacer(key) +
					'</label><div class="col-sm-7"><p class="form-control-static">' +
					identity[key] +
					'</p></div>'
			}
		}
		hr.innerHTML = buf
	}
}
xhr.open('GET', 'https://identigen.herokuapp.com', true);
xhr.send(null);
}
regen()


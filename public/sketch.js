var loginInput, passwordInput, buttonLogin

function setup() {
	createCanvas(windowWidth, windowHeight);
	background(0);

	loginInput = createInput('');
	passwordInput = createInput('', 'password');

	buttonLogin = createButton('Enviar');
	buttonLogin.mousePressed(login);

	loginInput.position(windowWidth/2 - loginInput.width/2, windowHeight/2 + loginInput.height/2);
	passwordInput.position(windowWidth/2 - passwordInput.width/2, windowHeight/2 + passwordInput.height + loginInput.height/2);
	buttonLogin.position(windowWidth/2 - buttonLogin.width/2, windowHeight/2 + buttonLogin.height + passwordInput.height + loginInput.height/2);
}

function draw() {

}

function login() {

	var playerData = {
		"login" : loginInput.value(),
		"password" : passwordInput.value()
	};

	console.log(playerData);

	httpPost('/login', playerData, 'json');
}
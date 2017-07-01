var socket;
var emailInput, passwordInput, buttonLogin;

function setup() {
	createCanvas(windowWidth, windowHeight);
	background(0);


/* Necessário arrumar isso de uma forma melhor */
/* Apenas para teste de servidor */
	emailInput = createInput('');
	passwordInput = createInput('', 'password');

	buttonLogin = createButton('Enviar');
	buttonLogin.mousePressed(login);

	buttonRegister = createButton('Registrar');
	buttonRegister.mousePressed(register);

	emailInput.position(windowWidth/2 - emailInput.width/2, windowHeight/2 + emailInput.height/2);
	passwordInput.position(windowWidth/2 - passwordInput.width/2, windowHeight/2 + passwordInput.height + emailInput.height/2);
	buttonLogin.position(windowWidth/2 - buttonLogin.width/2, windowHeight/2 + buttonLogin.height + passwordInput.height + emailInput.height/2);
	buttonRegister.position(windowWidth/2 - buttonRegister.width/2, windowHeight/2 + buttonLogin.height + passwordInput.height + buttonLogin.height+ emailInput.height/2);

}

function draw() {

}

function login() {

	var accountData = {
		"email" : emailInput.value(),
		"password" : passwordInput.value()
	};

	console.log(accountData);

	httpPost('/login', accountData, 'json');
}

function register() {

	var accountData = {
		"email" : emailInput.value(),
		"password" : passwordInput.value()
	};

	console.log(accountData);

	httpPost('/register', accountData, 'json');	
}
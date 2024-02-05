<script>
	import Keypad from "./keypad.svelte";
	let entrada = ""
	let total = 0
</script>

<main>
	<h1>Calculadora basica</h1>

	<div class="display">
		<span>{entrada}</span>
	</div>
<div class="input-pad">

	<Keypad tipoBoton ="borrar">C</Keypad>
	<Keypad tipoBoton ="borrar">CE</Keypad>
	<Keypad tamanio={2} tipoBoton ="historial">Hist</Keypad>

	<Keypad>7</Keypad>
	<Keypad>8</Keypad>
	<Keypad>9</Keypad>
	<Keypad tipoBoton ="operador">/</Keypad>

	<Keypad>4</Keypad>
	<Keypad>5</Keypad>
	<Keypad>6</Keypad>
	<Keypad tipoBoton ="operador">*</Keypad>


	<Keypad>1</Keypad>
	<Keypad>2</Keypad>
	<Keypad>3</Keypad>
	<Keypad tipoBoton ="operador">-</Keypad>

	<Keypad>0</Keypad>
	<Keypad>.</Keypad>
	<Keypad tipoBoton ="operador">=</Keypad>
	<Keypad tipoBoton ="operador">+</Keypad>
	
</div>
	
</main>

<style>
.display {
	max-width: 440px;
	margin: auto;
    margin-bottom: 10px;
	background-color: rgb(108, 175, 125);
		border: 1px solid #000000;
		padding: 31px;
		margin-bottom: 10px;
	}

	main {
		text-align: center;
		padding: 1em;
		border-radius: 40px;
		margin: 0 auto;
		background-color: aqua;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}
   .input-pad {
	max-width: 500px;
	margin: auto;
	display: grid;
	grid-template-columns: repeat(4,1fr);
	grid-template-rows: repeat(5,1fr);
	gap:10px;
   }

</style>
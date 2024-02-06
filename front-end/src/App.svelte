<script>
	import Keypad from "./keypad.svelte";
	let entrada = "";
	let entrada2 ="";
	let operandoActual = ""; 
	let operacion ="";
	let b=0;
	let total = 0;

	function agregarDigito(digito) {
		console.log(digito)
    if (entrada.length < 2) {
      entrada += digito;
	  console.log(entrada)
    }
	if (b === 1){
		if (entrada2.length < 2) {
			entrada2 += 1;
            entrada += digito;
	}
  }
}
  function agregarOperador(operador) {
    entrada+=operador;
	b=1;
    console.log(entrada);
    operacion = operador;
}

  async function realizarCalculo(){
	console.log(entrada)
	console.log(operacion)
	let operandos = entrada.split(operacion);
	let ope1=operandos[0];
	let ope2=operandos[1];

	if (ope1 !== "" && ope2 !== "" && operacion !== ""){
	const data = {
		op1: parseInt(ope1),
        op2: parseInt(ope2),
        op: operacion,
	}
	console.log(data)
  try {
        const response = await fetch('http://localhost:8000/calcular', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(data),
        });
		if (response.ok) {
          const result = await response.json();
          total = result.res;
          entrada = total.toString(); // Actualiza la entrada con el resultado
        } else {
          console.error('Error en la solicitud:', response.statusText);
        }
} catch (error) {
        console.error('Error en la solicitud:', error.message);
      }}}

</script>

<main>
	<h1>Calculadora basica</h1>

	<div class="display">
		<span>{entrada}</span>
	</div>
<div class="input-pad">
	<Keypad tipoBoton ="borrar" on:click={() => entrada = ""}>C</Keypad>
	<Keypad tipoBoton ="borrar" on:click={() => { entrada = ""; operandoActual = ""; operacion = ""; total = 0; }}>CE</Keypad>
	<Keypad tamanio={2} tipoBoton ="historial" on:click={() => alert("No implementado")}>Hist</Keypad>

<Keypad clicked={() => agregarDigito(7)}>7</Keypad>
<Keypad clicked={() => agregarDigito(8)}>8</Keypad>
<Keypad clicked={() => agregarDigito(9)}>9</Keypad>
<Keypad clicked={() => agregarOperador("/")}>/</Keypad>

	<Keypad clicked={() => agregarDigito(4)}>4</Keypad>
<Keypad clicked={() => agregarDigito(5)}>5</Keypad>
<Keypad clicked={() => agregarDigito(6)}>6</Keypad>
<Keypad clicked={() => agregarOperador("*")}>*</Keypad>


	<Keypad clicked={() => agregarDigito(1)}>1</Keypad>
<Keypad clicked={() => agregarDigito(2)}>2</Keypad>
<Keypad clicked={() => agregarDigito(3)}>3</Keypad>
<Keypad clicked={() => agregarOperador("-")}>-</Keypad>

	<Keypad>0</Keypad>
	<Keypad>.</Keypad>
	<Keypad clicked={realizarCalculo} tipoBoton ="operador">=</Keypad>
	<Keypad clicked={() => agregarOperador("+")} tipoBoton ="operador">+</Keypad>
	
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
<script>
	import Keypad from "./keypad.svelte";
	let entrada = "";
	let entrada2 ="";
	let operacion ="";
	let b=0;
    let total=0;
	let historial = [];

	function agregarDigito(digito) {
		console.log(digito)
    if (entrada.length < 2 ) {
      entrada += digito;
	  console.log(entrada)
    }
	if (b === 1 && entrada2.length < 2) {
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
		  historial.push(data);
		} else {
          console.error('Error en la solicitud:', response.statusText);
        }
} catch (error) {
        console.error('Error en la solicitud:', error.message);
      }}}
	  const borrar = () => {
		entrada = "";
		entrada2= "";
		b=0;
	}
	async function obtenerHistorial() {
    try {
      const response = await fetch('http://localhost:8000/historial');
      if (response.ok) {
        historial = await response.json();
        console.log('Historial:', historial);
      } else {
        console.error('Error al obtener el historial:', response.statusText);
      }
    } catch (error) {
      console.error('Error al obtener el historial:', error.message);
    }
  }
</script>

<main>
	<h1>Calculadora basica</h1>

	<div class="display">
		<span>{entrada}</span>
	</div>
	<div class="historial">
        {#each historial as operacion, i}
            <div>{operacion.op1} {operacion.op} {operacion.op2} = {operacion.res}</div>
        {/each}
    </div>
<div class="input-pad">
	<Keypad tamanio={2} tipoBoton ="borrar" clicked={borrar}>C</Keypad>
	<Keypad tamanio={2} tipoBoton ="historial" clicked={obtenerHistorial}>Hist</Keypad>

<Keypad clicked={() => agregarDigito(7)}>7</Keypad>
<Keypad clicked={() => agregarDigito(8)}>8</Keypad>
<Keypad clicked={() => agregarDigito(9)}>9</Keypad>
<Keypad clicked={() => agregarOperador("/")} tipoBoton ="operador">/</Keypad>

	<Keypad clicked={() => agregarDigito(4)}>4</Keypad>
<Keypad clicked={() => agregarDigito(5)}>5</Keypad>
<Keypad clicked={() => agregarDigito(6)}>6</Keypad>
<Keypad clicked={() => agregarOperador("*")} tipoBoton ="operador">*</Keypad>


	<Keypad clicked={() => agregarDigito(1)}>1</Keypad>
<Keypad clicked={() => agregarDigito(2)}>2</Keypad>
<Keypad clicked={() => agregarDigito(3)}>3</Keypad>
<Keypad clicked={() => agregarOperador("-")} tipoBoton ="operador">-</Keypad>

<Keypad clicked={() => agregarDigito(0)}>0</Keypad>
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
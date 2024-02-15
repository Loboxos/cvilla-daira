<script>
 export let tamanio = 1;
 export let tipoBoton = 'numero';
 export let clicked = () => {
    console.log("click");
 };
 
 function getColor (){
    switch (tipoBoton){
        case "numero":
            return "bg-gray-600";
        case "operador":
            return "bg-gray-400";
        case "borrar":
            return "bg-red-600";
        case "historial":
            return "bg-black";
        default:
            return "bg-white";
    }
  }
function handleClick() {
		clicked();
	}
</script>
<button class="key-pad {tamanio === 2 ? 'col-span-2' : ''} {getColor()} text-white py-2 px-4 rounded" on:click={handleClick}>
    <slot></slot>
  </button>

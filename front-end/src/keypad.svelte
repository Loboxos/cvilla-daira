<script>
 export let tamanio = 1;
 export let tipoBoton = 'numero';
 export let clicked = () => {
    console.log("click");
 };
 
function getColor (){
    switch (tipoBoton){
        case "numero":
            return "#515050";
        case "operador":
            return "#7B7676"
        case "borrar":
            return "#EA3B3B"
        case "historial":
            return "#1F201F"
        default:
         return "#fff";
    }
}
function handleClick() {
		clicked();
	}
</script>
<button class = "key-pad" style="grid-column:span {tamanio};background-color:{getColor()}; color:white" on:click={handleClick}>
<slot></slot>
</button>

<style>
    .key-pad {
      border: 1px solid #515050;
      padding: 22px 20px;
      margin: 0;
      font-size: 16px;
    }

    
</style>


export default function PaginaAdmin(){

    return (<div className="paginaAdmin">


            <form>
                <input type="text" placeholder="Nombre Hotel"></input>
                <input type="text" placeholder="Habitaciones Disponibles"></input>
                <input type="text" placeholder="Precio"></input>
                <input type="text" placeholder="atributo"></input>
                <input type="text" placeholder="atributo"></input>
                <button>Cargar Hotel</button>
            </form>

            <form>
                <button>Ver Reservas</button>
            </form>

    </div>
    )

}
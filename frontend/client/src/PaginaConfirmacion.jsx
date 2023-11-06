import React from "react";
import { Link } from "react-router-dom";
    
    
    const navbarLinks = [

    { url: "/", title: "Reservar" },
    { url: "/misreservas", title: "Mis Reservas" },
    { url: "/login", title: "Login" },
      ];
      
      
      
      const ConfirmacionReserva = () => {
        return (
          <div className="confirmacion-reserva">
            <h2>Felicitaciones, su reserva ha sido exitosa</h2>
            <Link to="/">
              <button>Volver al menu principal</button>
            </Link>
          </div>
        );
      };
      
      export default ConfirmacionReserva;
      

import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import PaginaReserva from './PaginaReserva';
import PaginaLogin from './PaginaLogin';
import PaginaRegistro from './PaginaRegistro';
import PaginaAdmin from './PaginaAdmin';
import PaginaComandos from './PaginaComandos';
import PaginaConfirmacion from './PaginaConfirmacion';
import MisReservas from './MisReservas';
import PaginaRechazo from './PaginaRechazo';
import PaginaDisponibilidad from './PaginaDisponibilidad';
import PaginaInfohoteles from './PaginaInfohoteles';

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
            <Routes>
                <Route index element={<PaginaReserva />} />
                <Route path="/login" element={<PaginaLogin/> } />
                <Route path="/register" element={<PaginaRegistro/> } />
                <Route path="/admin" element={<PaginaAdmin/> } />
                <Route path="/comandos" element={<PaginaComandos/> } />
                <Route path="/misreservas" element={<MisReservas/> } />
                <Route path="/confirmacion" element={<PaginaConfirmacion/> } />
                <Route path="/rechazo" element={<PaginaRechazo/> } />
                <Route path="/disponibilidad" element={<PaginaDisponibilidad/> } />
                <Route path="/infohoteles" element={<PaginaInfohoteles/> } />
                <Route path="/rechazo" element={<PaginaRechazo/> } />
                <Route path="/*" element={<h1>err 404: This route doesn't exists</h1>} />
            </Routes>
        </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
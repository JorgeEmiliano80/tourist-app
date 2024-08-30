// Función para cargar contenido dinámicamente
async function cargarContenidoDinamico(url, elementoDestino) {
    try {
        const respuesta = await fetch(url);
        if (!respuesta.ok) {
            throw new Error('Error en la respuesta del servidor');
        }
        const datos = await respuesta.json();
        renderizarContenido(datos, elementoDestino);
    } catch (error) {
        console.error('Error al cargar el contenido:', error);
        elementoDestino.innerHTML = '<p>Error al cargar el contenido. Por favor, intente de nuevo más tarde.</p>';
    }
}

// Función para renderizar el contenido
function renderizarContenido(datos, elemento) {
    let contenidoHTML = '';
    datos.forEach(item => {
        contenidoHTML += `
            <div class="item">
                <h3>${item.titulo}</h3>
                <p>${item.descripcion}</p>
                <a href="${item.enlace}" class="btn btn-primary">Ver más</a>
            </div>
        `;
    });
    elemento.innerHTML = contenidoHTML;
}

// Función para manejar la búsqueda
async function manejarBusqueda(event) {
    event.preventDefault();
    const terminoBusqueda = document.getElementById('busqueda-input').value;
    const resultadosElemento = document.getElementById('resultados-busqueda');
    resultadosElemento.innerHTML = '<p>Buscando...</p>';
    
    try {
        const respuesta = await fetch(`/api/buscar?q=${encodeURIComponent(terminoBusqueda)}`);
        if (!respuesta.ok) {
            throw new Error('Error en la búsqueda');
        }
        const resultados = await respuesta.json();
        renderizarResultadosBusqueda(resultados, resultadosElemento);
    } catch (error) {
        console.error('Error en la búsqueda:', error);
        resultadosElemento.innerHTML = '<p>Error al realizar la búsqueda. Por favor, intente de nuevo.</p>';
    }
}

// Función para renderizar los resultados de la búsqueda
function renderizarResultadosBusqueda(resultados, elemento) {
    if (resultados.length === 0) {
        elemento.innerHTML = '<p>No se encontraron resultados.</p>';
        return;
    }

    let contenidoHTML = '<ul>';
    resultados.forEach(resultado => {
        contenidoHTML += `<li><a href="${resultado.url}">${resultado.titulo}</a></li>`;
    });
    contenidoHTML += '</ul>';
    elemento.innerHTML = contenidoHTML;
}

// Función para cambiar el idioma
async function cambiarIdioma(idioma) {
    try {
        const respuesta = await fetch(`/api/cambiar-idioma?lang=${idioma}`, { method: 'POST' });
        if (!respuesta.ok) {
            throw new Error('Error al cambiar el idioma');
        }
        const resultado = await respuesta.json();
        if (resultado.success) {
            location.reload(); // Recargar la página para aplicar los cambios de idioma
        } else {
            throw new Error('No se pudo cambiar el idioma');
        }
    } catch (error) {
        console.error('Error al cambiar el idioma:', error);
        alert('No se pudo cambiar el idioma. Por favor, intente de nuevo.');
    }
}

// Función para manejar el menú móvil
function toggleMenuMovil() {
    const menu = document.querySelector('.menu-movil');
    menu.classList.toggle('activo');
    
    // Añadir/quitar clase al body para prevenir el scroll
    document.body.classList.toggle('menu-abierto');
}

// Función para manejar las preguntas frecuentes
function togglePreguntaFrecuente(elemento) {
    elemento.classList.toggle('activa');
    const respuesta = elemento.querySelector('.faq-respuesta');
    if (respuesta.style.maxHeight) {
        respuesta.style.maxHeight = null;
    } else {
        respuesta.style.maxHeight = respuesta.scrollHeight + "px";
    }
}

// Función para el desplazamiento suave
function desplazamientoSuave(evento) {
    evento.preventDefault();
    const destino = document.querySelector(evento.target.getAttribute('href'));
    destino.scrollIntoView({
        behavior: 'smooth',
        block: 'start'
    });
}

// Función para cargar más contenido
async function cargarMasContenido() {
    const contenedorContenido = document.getElementById('contenedor-contenido');
    const botonCargarMas = document.getElementById('cargar-mas');
    const paginaActual = parseInt(botonCargarMas.dataset.pagina) || 1;
    
    try {
        botonCargarMas.textContent = 'Cargando...';
        botonCargarMas.disabled = true;
        
        const respuesta = await fetch(`/api/contenido?pagina=${paginaActual + 1}`);
        if (!respuesta.ok) {
            throw new Error('Error al cargar más contenido');
        }
        const nuevoContenido = await respuesta.json();
        
        if (nuevoContenido.length > 0) {
            renderizarContenido(nuevoContenido, contenedorContenido);
            botonCargarMas.dataset.pagina = paginaActual + 1;
            botonCargarMas.textContent = 'Cargar más';
            botonCargarMas.disabled = false;
        } else {
            botonCargarMas.textContent = 'No hay más contenido';
            botonCargarMas.disabled = true;
        }
    } catch (error) {
        console.error('Error al cargar más contenido:', error);
        botonCargarMas.textContent = 'Error al cargar. Intentar de nuevo';
        botonCargarMas.disabled = false;
    }
}

// Inicialización cuando el DOM está completamente cargado
document.addEventListener('DOMContentLoaded', function() {
    // Cargar contenido inicial
    const contenedorPrincipal = document.getElementById('contenedor-principal');
    if (contenedorPrincipal) {
        cargarContenidoDinamico('/api/contenido-inicial', contenedorPrincipal);
    }

    // Manejar el formulario de búsqueda
    const formularioBusqueda = document.getElementById('formulario-busqueda');
    if (formularioBusqueda) {
        formularioBusqueda.addEventListener('submit', manejarBusqueda);
    }

    // Manejar el cambio de idioma
    const selectorIdioma = document.getElementById('selector-idioma');
    if (selectorIdioma) {
        selectorIdioma.addEventListener('change', function() {
            cambiarIdioma(this.value);
        });
    }

    // Manejar el menú móvil
    const botonMenuMovil = document.querySelector('.boton-menu-movil');
    if (botonMenuMovil) {
        botonMenuMovil.addEventListener('click', toggleMenuMovil);
    }

    // Manejar las preguntas frecuentes
    const preguntasFrecuentes = document.querySelectorAll('.faq-pregunta');
    preguntasFrecuentes.forEach(pregunta => {
        pregunta.addEventListener('click', function() {
            togglePreguntaFrecuente(this.parentElement);
        });
    });

    // Manejar el desplazamiento suave para los enlaces internos
    const enlacesInternos = document.querySelectorAll('a[href^="#"]');
    enlacesInternos.forEach(enlace => {
        enlace.addEventListener('click', desplazamientoSuave);
    });

    // Manejar el botón "Cargar más"
    const botonCargarMas = document.getElementById('cargar-mas');
    if (botonCargarMas) {
        botonCargarMas.addEventListener('click', cargarMasContenido);
    }

    console.log('DOM completamente cargado y funcionalidades inicializadas');
});


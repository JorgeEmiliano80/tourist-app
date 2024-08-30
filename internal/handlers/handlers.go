package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// Estructura para almacenar las plantillas
type TemplateHandler struct {
	templates *template.Template
}

// Función para crear un nuevo TemplateHandler
func NewTemplateHandler(templateDir string) (*TemplateHandler, error) {
	templates, err := template.ParseGlob(templateDir + "/*.html")
	if err != nil {
		return nil, err
	}
	return &TemplateHandler{templates: templates}, nil
}

// Manejador para la página de inicio
func (th *TemplateHandler) Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title":       "Bienvenido a nuestra agencia de viajes",
		"Description": "Descubre los mejores tours y experiencias",
	}
	th.templates.ExecuteTemplate(w, "home.html", data)
}

// Manejador para la página de tours
func (th *TemplateHandler) Tours(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la lógica para obtener los tours de la base de datos
	tours := []map[string]interface{}{
		{"Name": "Tour por la ciudad", "Duration": "3 horas", "Price": 50},
		{"Name": "Excursión a la montaña", "Duration": "8 horas", "Price": 100},
	}
	data := map[string]interface{}{
		"Title": "Nuestros Tours",
		"Tours": tours,
	}
	th.templates.ExecuteTemplate(w, "tours.html", data)
}

// Manejador para la página de pagos
func (th *TemplateHandler) Payments(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Procesar el pago
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}
		amount := r.FormValue("amount")
		currency := r.FormValue("currency")
		// Aquí iría la lógica para procesar el pago
		log.Printf("Procesando pago: %s %s", amount, currency)
		response := map[string]string{"status": "success", "message": "Pago procesado correctamente"}
		json.NewEncoder(w).Encode(response)
	} else {
		data := map[string]interface{}{
			"Title": "Realizar Pago",
		}
		th.templates.ExecuteTemplate(w, "payments.html", data)
	}
}

// Manejador para la página de preguntas frecuentes
func (th *TemplateHandler) FAQ(w http.ResponseWriter, r *http.Request) {
	faqs := []map[string]string{
		{"Question": "¿Cuál es la política de cancelación?", "Answer": "Puedes cancelar hasta 24 horas antes del tour."},
		{"Question": "¿Los tours incluyen comidas?", "Answer": "Depende del tour específico. Revisa los detalles de cada tour."},
	}
	data := map[string]interface{}{
		"Title": "Preguntas Frecuentes",
		"FAQs":  faqs,
	}
	th.templates.ExecuteTemplate(w, "faq.html", data)
}

// Manejador para la página de contacto
func (th *TemplateHandler) Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Procesar el formulario de contacto
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")
		// Aquí iría la lógica para guardar el mensaje o enviar un correo
		log.Printf("Mensaje recibido de %s (%s): %s", name, email, message)
		response := map[string]string{"status": "success", "message": "Mensaje enviado correctamente"}
		json.NewEncoder(w).Encode(response)
	} else {
		data := map[string]interface{}{
			"Title": "Contáctanos",
		}
		th.templates.ExecuteTemplate(w, "contact.html", data)
	}
}

// Manejador para la página de error 404
func (th *TemplateHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	data := map[string]interface{}{
		"Title": "Página no encontrada",
	}
	th.templates.ExecuteTemplate(w, "404.html", data)
}

import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { ThemeProvider, CssBaseline } from '@mui/material';
import { AuthProvider } from './context/AuthContext';
import { ClientProvider } from './context/ClientContext';
import Header from '../src/components/Header';
import Footer from '../src/components/Footer';
import HomePage from '../src/pages/HomePage';
import TourPage from '../src/pages/TourPage';
import BookingPage from '../src/pages/BookingPages';
import PaymentPage from '../src/pages/PaymentPage';
import ClientPage from '../src/pages/ClientPage';
import theme from '../src/styles/theme';
import '../src/styles/App.css';

function App() {
    return (
        <ThemeProvider theme={theme}>
            <CssBaseline />
            <AuthProvider>
                <ClientProvider>
                    <Router>
                        <div className="App">
                            <Header />
                            <main>
                                <Routes>
                                    <Route path="/" element={<HomePage />} />
                                    <Route path="/tours/:id" element={<TourPage />} />
                                    <Route path="/bookings" element={<BookingPage />} />
                                    <Route path="/payments" element={<PaymentPage />} />
                                    <Route path="/clients" element={<ClientPage />} />
                                </Routes>
                            </main>
                            <Footer />
                        </div>
                    </Router>
                </ClientProvider>
            </AuthProvider>
        </ThemeProvider>
    );
}

export default App;
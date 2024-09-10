import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { ThemeProvider, CssBaseline } from '@material-ui/core';
import { AuthProvider } from './context/AuthContext';
import { ClientProvider } from './context/ClientContext';
import Header from './components/Header';
import Footer from './components/Footer';
import HomePage from './pages/HomePage';
import TourPage from './pages/TourPage';
import BookingPage from './pages/BookingPage';
import PaymentPage from './pages/PaymentPage';
import ClientPage from './pages/ClientPage';
import theme from './styles/theme';
import './styles/App.css';

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
                                <Switch>
                                    <Route exact path="/" component={HomePage} />
                                    <Route path="/tours/:id" component={TourPage} />
                                    <Route path="/bookings" component={BookingPage} />
                                    <Route path="/payments" component={PaymentPage} />
                                    <Route path="/clients" component={ClientPage} />
                                </Switch>
                            </main>
                            <Footer />
                        </div>
                    </Router>
                </ClientProvider>
            </AuthProvider>
        </ThemeProvider>
    );
};

export default App;

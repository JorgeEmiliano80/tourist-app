import React, { useState } from 'react';
import { Container, Typography, Button } from '@mui/material';
import ClientList from '../components/ClientList';
import { useClientContext } from '../context/ClientContext';
import { createClient, updateClient } from '../services/clientService';

const ClienPage = () => {
    const { clients, fetchClients } = useClientContext();
    const [editingClient, setEditingClient] = useState(null);

    const handleCreateClient = async (clientData) => {
        try {
            await createClient(clientData);
            fetchClients();
        } catch (error) {
            console.error('Error creating client: ', error);
        }
    };

    const handleUpdateClient = async (id, clientData) => {
        try {
            await updateClient(id, clientData);
            fetchClients();
            setEditingClient(null);
        } catch (error) {
            console.error('Error updating client: ', error);
        }
    };

    return (
        <Container>
            <Typography variant="h2" gutterBottom>
                Clients
            </Typography>
            <Button
                variant="contained"
                color="rpimary"
                onClick={() => setEditingClient({})}
            >
                Add New Client
            </Button>
            {editingClient && (
                <ClientForm
                    client={editingClient}
                    onSubmit={editingClient.id ? handleUpdateClient : handleCreateClient}
                    onCancel={() => setEditingClient(null)}
                />
            )}
            <ClientList
                clients={clients}
                onEdit={setEditingClient}
            />
        </Container>
    );
};

export default ClientPage;

import React from 'react';
import { List, ListItem, ListItemText, IconButton, Stack } from '@mui/material';
import { Delete, Edit } from '@mui/icons-material';
import { useClientContext } from '../context/ClientContext';

const ClientList = ({ onEdit }) => {
    const { clients, fetchClients, deleteClient } = useClientContext();

    const handleDelete = async (id) => {
        try {
            await deleteClient(id);
            fetchClients();
        } catch (error) {
            console.error('Error deleting client: ', error);
        }
    };

    return (
        <List>
            {clients.map((client) => (
                <ListItem
                    key={client.id}
                    secondaryAction={
                        <Stack direction="row" spacing={1}>
                            <IconButton edge="end" aria-label="edit" onClick={() => onEdit(client)}>
                                <Edit />
                            </IconButton>
                            <IconButton edge="end" aria-label="delete" onClick={() => handleDelete(client.id)}>
                                <Delete />
                            </IconButton>
                        </Stack>
                    }
                >
                    <ListItemText primary={client.name} secondary={client.email} />
                </ListItem>
            ))}
        </List>
    );
};

export default ClientList;
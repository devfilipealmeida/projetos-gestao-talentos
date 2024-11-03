import { useState } from "react";
import { useQuery, useMutation, useQueryClient } from "react-query";
import { Container, Typography, Button, List, ListItem, ListItemText, ListItemButton, Box, Chip, TextField, Select, MenuItem } from "@mui/material";
import { Add } from "@mui/icons-material";
import { useNavigate } from "react-router-dom";
import CandidatoForm from "../components/CandidatoForm/candidatoForm.component";
import api from "../api/axios";

const fetchCandidates = async ({ queryKey }: any) => {
  const [, { search, status }] = queryKey;
  const response = await api.get("/candidatos", {
    params: {
      search,
      status,
    },
  });
  return response.data;
};

const Candidatos = () => {
  const queryClient = useQueryClient();
  const navigate = useNavigate();
  const [isModalOpen, setModalOpen] = useState(false);
  const [selectedCandidateId, setSelectedCandidateId] = useState<string | undefined>(undefined);
  const [search, setSearch] = useState('');
  const [statusFilter, setStatusFilter] = useState('');

  const { data: candidates } = useQuery(
    ["candidatos", { search, status: statusFilter }],
    fetchCandidates,
    {
      refetchOnWindowFocus: false,
    }
  );

  const mutation = useMutation(
    (data: { data: any; id?: string }) => {
      if (data.id) {
        return api.put(`/candidatos/${data.id}`, data.data);
      } else {
        return api.post("/candidatos", data.data);
      }
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("candidatos");
        handleCloseModal();
      },
      onError: (error) => {
        console.error("Erro ao salvar candidato:", error);
      },
    }
  );

  const handleOpenModal = (candidateId?: string) => {
    setSelectedCandidateId(candidateId);
    setModalOpen(true);
  };

  const handleCloseModal = () => {
    setModalOpen(false);
    setSelectedCandidateId(undefined);
  };

  const handleSaveCandidate = (data: any, candidateId?: string) => {
    mutation.mutate({ data, id: candidateId });
  };

  const handleDeleteCandidate = () => {
    queryClient.invalidateQueries("candidatos");
  };

  const handleLogoff = () => {
    localStorage.removeItem('token');
    navigate('/login');
  };

  return (
    <Container maxWidth="lg" sx={{ mt: 4, textAlign: "center" }}>
      <Typography variant="h5" fontWeight="bold" mb={1}>
        Bem-vindo 👋
      </Typography>
      <Typography variant="body1" color="textSecondary" mb={3}>
        Gerenciamento de Recursos Humanos.
      </Typography>

      <Box display="flex" justifyContent="flex-end" mb={2}>
        <Button variant="outlined" color="secondary" onClick={handleLogoff}>
          Logoff
        </Button>
      </Box>

      <Box display="flex" justifyContent="flex-start" alignItems="center" gap={2} mb={2} sx={{ maxWidth: '80%', mx: 'auto' }}>
        <Select
          value={statusFilter}
          onChange={(e) => setStatusFilter(e.target.value)}
          displayEmpty
          variant="outlined"
          size="small"
          sx={{ minWidth: 120 }}
        >
          <MenuItem value="">Filtrar Status</MenuItem>
          <MenuItem value="Ativo">Ativo</MenuItem>
          <MenuItem value="Inativo">Inativo</MenuItem>
        </Select>

        <TextField
          label="Pesquisar candidato"
          variant="outlined"
          size="small"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
        />
        <Button
          variant="contained"
          startIcon={<Add />}
          size="small"
          sx={{ height: '40px' }}
          onClick={() => handleOpenModal()}
        >
          Novo
        </Button>
      </Box>

      <Box display="flex" justifyContent="center" p={1} sx={{ borderBottom: '1px solid #e0e0e0', backgroundColor: '#f5f5f5', maxWidth: '80%', mx: 'auto' }}>
        <Typography variant="subtitle2" fontWeight="bold" width="25%" textAlign="center">
          Nome
        </Typography>
        <Typography variant="subtitle2" fontWeight="bold" width="20%" textAlign="center">
          Status
        </Typography>
        <Typography variant="subtitle2" fontWeight="bold" width="55%" textAlign="center">
          Habilidades
        </Typography>
      </Box>
      
      <List sx={{ maxWidth: "80%", mx: "auto" }}>
        {candidates?.map((candidate: any) => (
          <ListItem
            key={candidate.id}
            sx={{ display: "flex", justifyContent: "space-between", p: 1, borderBottom: "1px solid #e0e0e0" }}
          >
            <ListItemButton onClick={() => handleOpenModal(candidate.id)}>
              <ListItemText primary={candidate.nome} sx={{ width: "25%", textAlign: "center" }} />
              <ListItemText primary={candidate.status} sx={{ width: "20%", textAlign: "center" }} />
              <Box display="flex" flexWrap="wrap" gap={0.5} sx={{ width: "55%", justifyContent: "center" }}>
                {candidate.habilidades.map((habilidade: string, index: number) => (
                  <Chip key={index} label={habilidade} color="primary" variant="outlined" size="small" />
                ))}
              </Box>
            </ListItemButton>
          </ListItem>
        ))}
      </List>

      <CandidatoForm
        isModalOpen={isModalOpen}
        handleCloseModal={handleCloseModal}
        onSave={handleSaveCandidate}
        onDelete={handleDeleteCandidate}
        isLoading={mutation.isLoading}
        candidateId={selectedCandidateId}
      />
    </Container>
  );
};

export default Candidatos;
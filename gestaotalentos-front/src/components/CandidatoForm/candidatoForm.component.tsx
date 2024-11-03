import { useState, useEffect } from "react";
import { Close } from "@mui/icons-material";
import { Box, Button, Chip, IconButton, MenuItem, Modal, Select, TextField, Typography } from "@mui/material";
import { useForm, Controller } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import api from "../../api/axios";
import ConfirmDeleteDialog from "../DeleteDialog/delete.component";

const candidateSchema = z.object({
  nome: z.string().nonempty("Nome é obrigatório"),
  status: z.enum(["Ativo", "Inativo"], { invalid_type_error: "Status é obrigatório" }),
  habilidades: z.array(z.string()).min(1, "Ao menos uma habilidade é obrigatória"),
});

type CandidateFormValues = z.infer<typeof candidateSchema>;

interface CandidatoFormProps {
  isModalOpen: boolean;
  handleCloseModal: (resetForm: () => void) => void;
  onSave: (data: CandidateFormValues, id?: string) => void;
  onDelete: () => void;
  isLoading: boolean;
  candidateId?: string;
}

const CandidatoForm: React.FC<CandidatoFormProps> = ({ isModalOpen, handleCloseModal, onSave, onDelete, isLoading, candidateId }) => {
  const { control, handleSubmit, setValue, watch, reset, formState: { errors } } = useForm<CandidateFormValues>({
    resolver: zodResolver(candidateSchema),
    defaultValues: { nome: "", status: "Ativo", habilidades: [] },
    mode: 'onChange',
  });

  const [habilidadeInput, setHabilidadeInput] = useState('');
  const [confirmDeleteOpen, setConfirmDeleteOpen] = useState(false);
  const habilidades = watch("habilidades");

  useEffect(() => {
    if (candidateId) {
      const fetchCandidate = async () => {
        try {
          const response = await api.get(`/candidatos/${candidateId}`);
          const candidateData = response.data;
          setValue("nome", candidateData.nome);
          setValue("status", candidateData.status);
          setValue("habilidades", candidateData.habilidades || []);
        } catch (error) {
          console.error("Erro ao buscar dados do candidato:", error);
        }
      };
      fetchCandidate();
    } else {
      reset();
    }
  }, [candidateId, setValue, reset]);

  const handleAddHabilidade = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter' && habilidadeInput.trim()) {
      setValue("habilidades", [...habilidades, habilidadeInput.trim()]);
      setHabilidadeInput('');
      event.preventDefault();
    }
  };

  const handleDeleteHabilidade = (habilidadeToDelete: string) => {
    setValue("habilidades", habilidades.filter((habilidade) => habilidade !== habilidadeToDelete));
  };

  const onSubmit = (data: CandidateFormValues) => {
    onSave(data, candidateId);
    handleCloseModal(reset);
    setHabilidadeInput('');
  };

  const handleDeleteCandidate = async () => {
    if (candidateId) {
      try {
        await api.delete(`/candidatos/${candidateId}`);
        onDelete();
        setConfirmDeleteOpen(false);
        handleCloseModal(reset);
      } catch (error) {
        console.error("Erro ao deletar candidato:", error);
      }
    }
  };

  const handleCloseAllModals = () => {
    setConfirmDeleteOpen(false);
    handleCloseModal(reset);
  };

  return (
    <Modal open={isModalOpen} onClose={() => handleCloseAllModals()}>
      <Box sx={{ position: 'absolute', top: '50%', left: '50%', transform: 'translate(-50%, -50%)', width: 400, bgcolor: 'background.paper', p: 4, borderRadius: 2, boxShadow: 24 }}>
        <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
          <Typography variant="h6">{candidateId ? "Editar Candidato" : "Adicionar Candidato"}</Typography>
          <IconButton onClick={() => handleCloseAllModals()}>
            <Close />
          </IconButton>
        </Box>
        <form onSubmit={handleSubmit(onSubmit)}>
          <Controller
            name="nome"
            control={control}
            render={({ field }) => (
              <TextField
                {...field}
                label="Nome"
                variant="outlined"
                fullWidth
                margin="normal"
                error={!!errors.nome}
                helperText={errors.nome?.message}
                autoComplete="off"
              />
            )}
          />
          <Controller
            name="status"
            control={control}
            render={({ field }) => (
              <Select
                {...field}
                displayEmpty
                variant="outlined"
                fullWidth
                error={!!errors.status}
              >
                <MenuItem value="">Selecionar Status</MenuItem>
                <MenuItem value="Ativo">Ativo</MenuItem>
                <MenuItem value="Inativo">Inativo</MenuItem>
              </Select>
            )}
          />
          {errors.status && (
            <Typography variant="caption" color="error">
              {errors.status.message}
            </Typography>
          )}
          <TextField
            label="Habilidades"
            variant="outlined"
            fullWidth
            margin="normal"
            value={habilidadeInput}
            onChange={(e) => setHabilidadeInput(e.target.value)}
            onKeyDown={handleAddHabilidade}
          />
          <Box display="flex" flexWrap="wrap" gap={1} mt={1}>
            {habilidades.map((habilidade, index) => (
              <Chip
                key={index}
                label={habilidade}
                onDelete={() => handleDeleteHabilidade(habilidade)}
                color="primary"
                variant="outlined"
                size="small"
              />
            ))}
          </Box>
          {errors.habilidades && (
            <Typography variant="caption" color="error">
              {errors.habilidades.message}
            </Typography>
          )}
          <Box display="flex" justifyContent="space-between" mt={2}>
            {candidateId && (
              <Button
                variant="outlined"
                color="error"
                onClick={() => setConfirmDeleteOpen(true)}
                sx={{ width: '48%' }}
              >
                Deletar
              </Button>
            )}
            <Button
              variant="contained"
              color="primary"
              type="submit"
              fullWidth={!candidateId}
              sx={{ width: candidateId ? '48%' : '100%' }}
              disabled={isLoading}
            >
              {isLoading ? 'Salvando...' : candidateId ? 'Atualizar' : 'Salvar'}
            </Button>
          </Box>
        </form>

        <ConfirmDeleteDialog
          open={confirmDeleteOpen}
          onClose={() => setConfirmDeleteOpen(false)}
          onConfirm={handleDeleteCandidate}
          title="Confirmar Exclusão"
          description="Tem certeza de que deseja excluir este candidato? Esta ação não pode ser desfeita."
        />
      </Box>
    </Modal>
  );
};

export default CandidatoForm;
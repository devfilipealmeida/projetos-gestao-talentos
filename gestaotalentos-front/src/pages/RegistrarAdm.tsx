import React from 'react';
import { useForm, Controller } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import { TextField, Button, Select, MenuItem, FormControl, InputLabel } from '@mui/material';
import { useMutation } from 'react-query';
import { useNavigate } from 'react-router-dom';
import api from '../api/axios';

const schema = z.object({
    nome: z.string().nonempty("Nome é obrigatório"),
    email: z.string().email("Email inválido").nonempty("Email é obrigatório"),
    senha: z.string().min(6, "Senha deve ter no mínimo 6 caracteres"),
    role: z.enum(["admin", "user"]).refine((val) => !!val, {
      message: "Selecione um tipo de usuário",
    }),
});  

type FormData = z.infer<typeof schema>;

const RegistrarAdm: React.FC = () => {
  const navigate = useNavigate();
  const { control, handleSubmit, formState: { errors } } = useForm<FormData>({
    resolver: zodResolver(schema),
  });

  const mutation = useMutation((data: FormData) =>
    api.post('/register', data)
  );

  const onSubmit = (data: FormData) => {
    mutation.mutate(data, {
      onSuccess: () => {
        alert("Usuário cadastrado com sucesso!");
        navigate('/login');
      },
      onError: (error) => {
        alert("Erro ao cadastrar usuário");
      },
    });
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} style={{ maxWidth: 400, margin: 'auto' }}>
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
          />
        )}
      />

      <Controller
        name="email"
        control={control}
        render={({ field }) => (
          <TextField
            {...field}
            label="Email"
            variant="outlined"
            fullWidth
            margin="normal"
            error={!!errors.email}
            helperText={errors.email?.message}
          />
        )}
      />

      <Controller
        name="senha"
        control={control}
        render={({ field }) => (
          <TextField
            {...field}
            label="Senha"
            type="password"
            variant="outlined"
            fullWidth
            margin="normal"
            error={!!errors.senha}
            helperText={errors.senha?.message}
          />
        )}
      />

      <Controller
        name="role"
        control={control}
        render={({ field }) => (
          <FormControl fullWidth margin="normal" variant="outlined">
            <InputLabel>Tipo de Usuário</InputLabel>
            <Select
              {...field}
              label="Tipo de Usuário"
              error={!!errors.role}
            >
              <MenuItem value="admin">Admin</MenuItem>
              <MenuItem value="user">User</MenuItem>
            </Select>
          </FormControl>
        )}
      />
      {errors.role && <p style={{ color: 'red' }}>{errors.role.message}</p>}

      <Button
        type="submit"
        variant="contained"
        color="primary"
        fullWidth
        disabled={mutation.isLoading}
        style={{ marginTop: 16 }}
      >
        {mutation.isLoading ? "Enviando..." : "Registrar"}
      </Button>
    </form>
  );
};

export default RegistrarAdm;
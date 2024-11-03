interface DecodedToken {
    exp: number;
    [key: string]: any;
  }
  
  export const decodeToken = (token: string): DecodedToken | null => {
    try {
      const payloadBase64 = token.split('.')[1];
      const decodedPayload = JSON.parse(atob(payloadBase64));
      return decodedPayload;
    } catch (error) {
      console.error("Erro ao decodificar o token:", error);
      return null;
    }
  };
  
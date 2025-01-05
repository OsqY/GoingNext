import Cookies from "js-cookie";
import { api } from "..";

export const filesService = {
  async sendFileToS3(file: File) {
    try {
      const token = Cookies.get("token");
      const formData = new FormData();
      formData.append("file", file);

      const response = await api.post<string>("/api/files/sends3", formData, {
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "multipart/form-data",
        },
      });

      return response.data;
    } catch (error) {
      console.error(error);
    }
  },
};

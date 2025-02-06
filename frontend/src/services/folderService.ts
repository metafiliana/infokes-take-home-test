import axios from 'axios';
import type { Folder } from '@/models/Folder';

export const folderService = {
  async getAllFolders(): Promise<Folder[]> {
    try {
      const response = await axios.get('/api/folders');
      return response.data.data; // Extract the data array
    } catch (error) {
      console.error('Error fetching folders:', error);
      return [];
    }
  },

  async getSubfolders(folderId: string): Promise<Folder[]> {
    try {
      const response = await axios.get(`/api/folders/${folderId}`);
      return response.data.data; // Extract the data array
    } catch (error) {
      console.error(`Error fetching subfolders for folder ${folderId}:`, error);
      return [];
    }
  }
};
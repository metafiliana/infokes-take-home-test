export interface File {
  id: string;
  name: string;
}

export interface Folder {
  id: string;
  name: string;
  parentId?: string;
  level: number;
  files?: File[];
}
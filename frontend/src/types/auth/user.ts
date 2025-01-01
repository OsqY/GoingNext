export interface UserInfo {
  ID: number;
  Email: string;
  Username: string;
  Password: string;
  RoleID: number;
  IsActive: boolean;
  ImageUrl: string | null;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string;
  CreatedBy: number;
  UpdatedBy: number;
  DeletedBy: number;
  RoleName: string;
}

export interface UpdateUserDTO {
  ID: number;
  Email: string;
  Username: string;
  Password: string;
  RoleID: number;
  ImageUrl: string | null;
}

export type RolesResponse = RoleDTO[];

export interface RoleDTO {
  ID: number;
  Name: string;
  Description: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string;
  CreatedBy: number;
  UpdatedBy: string;
  DeletedBy: number;
}

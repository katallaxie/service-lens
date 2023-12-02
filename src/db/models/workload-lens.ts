import {
  BelongsTo,
  Column,
  CreatedAt,
  DataType,
  DeletedAt,
  ForeignKey,
  HasMany,
  Max,
  Min,
  Model,
  NotEmpty,
  PrimaryKey,
  Table,
  UpdatedAt
} from 'sequelize-typescript'
import { Workload } from './workload'
import { Lens } from './lens'

export interface WorkloadLensAttributes {
  id: number
  lensId: string
  workloadId: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type WorkloadLensCreationAttributes = Omit<
  WorkloadLensAttributes,
  'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'workload-lens'
})
export class WorkloadLens extends Model<
  WorkloadLensAttributes,
  WorkloadLensCreationAttributes
> {
  @PrimaryKey
  @Column(DataType.INTEGER)
  id!: string

  @ForeignKey(() => Workload)
  @Column(DataType.UUIDV4)
  workloadId?: string

  @ForeignKey(() => Lens)
  @Column(DataType.UUIDV4)
  lensId?: string

  @CreatedAt
  @Column
  createdAt?: Date

  @UpdatedAt
  @Column
  updatedAt?: Date

  @DeletedAt
  @Column
  deletedAt?: Date
}
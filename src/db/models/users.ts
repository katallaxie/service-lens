import {
  Table,
  Model,
  Column,
  PrimaryKey,
  DataType,
  AutoIncrement,
  Unique
} from 'sequelize-typescript'

export interface UserAttributes {
  id: string
  email: string
  name: string
  emailVerified: string
  image?: string
}

export type UserCreationAttributes = Omit<UserAttributes, 'id'>

@Table({
  tableName: 'users',
  timestamps: false
})
export class User extends Model<UserAttributes, UserCreationAttributes> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.INTEGER)
  id!: string

  @Column
  name?: string

  @Unique
  @Column
  email!: string

  @Column
  emailVerified?: Date

  @Column
  image?: string
}

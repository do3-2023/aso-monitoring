import { Entity, Column, PrimaryGeneratedColumn } from 'typeorm';

@Entity()
export class Person {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ name: 'last_name' })
  lastName: string;

  @Column({ name: 'phone_number' })
  phoneNumber: string;

  @Column()
  location: string;
}

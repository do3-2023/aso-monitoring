import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Person } from './person.entity';
import { CreatePersonDto } from './create-person.dto';

@Injectable()
export class PersonService {
  constructor(
    @InjectRepository(Person)
    private personRepository: Repository<Person>,
  ) {}

  findAll(): Promise<Person[]> {
    return this.personRepository.find();
  }

  findOne(id: number): Promise<Person | null> {
    return this.personRepository.findOneBy({ id });
  }

  async remove(id: number): Promise<void> {
    await this.personRepository.delete(id);
  }

  async create(person: CreatePersonDto): Promise<Person> {
    return this.personRepository.save(person);
  }
}

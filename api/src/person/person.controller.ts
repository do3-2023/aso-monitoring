import { Body, Controller, Delete, Get, Param, Post } from '@nestjs/common';
import { PersonService } from './person.service';
import { Person } from './person.entity';
import { CreatePersonDto } from './create-person.dto';

@Controller('person')
export class PersonController {
  constructor(private personService: PersonService) {}

  @Get()
  findAll(): Promise<Person[]> {
    return this.personService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: number): Promise<Person | null> {
    return this.personService.findOne(id);
  }

  @Delete(':id')
  async remove(@Param('id') id: number): Promise<void> {
    await this.personService.remove(id);
  }

  @Post()
  async create(@Body() person: CreatePersonDto): Promise<Person> {
    return this.personService.create(person);
  }
}

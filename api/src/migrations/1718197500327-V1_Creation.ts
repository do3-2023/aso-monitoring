import { MigrationInterface, QueryRunner } from 'typeorm';

export class V1Creation1718197500327 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `
          CREATE TABLE person
          (
              id           SERIAL NOT NULL PRIMARY KEY,
              last_name    TEXT   NOT NULL,
              phone_number TEXT   NOT NULL,
              location     TEXT   NOT NULL
          );

          INSERT INTO person (last_name, phone_number, location)
          VALUES ('John', '0702030405', 'Marseille'),
                 ('Doe', '0603040506', 'Montpellier');
          `,
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query('DROP TABLE person;');
  }
}

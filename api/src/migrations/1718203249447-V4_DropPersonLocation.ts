import { MigrationInterface, QueryRunner } from 'typeorm';

export class V4DropPersonLocation1718203249447 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query('ALTER TABLE person DROP COLUMN location;');
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query('ALTER TABLE person ADD COLUMN location TEXT;');
  }
}

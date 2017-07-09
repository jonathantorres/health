<?php

namespace App\Console\Commands;

use Carbon\Carbon;
use Illuminate\Console\Command;
use Illuminate\Support\Facades\DB;

class ImportReadings extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'readings:import';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Import initial blood pressure readings.';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $this->info('Importing predefined readings...');

        DB::table('blood_pressures')->insert($this->readingsToImport());

        $this->info('Done.');
    }

    /**
     * Define and return the readings to import.
     *
     * @todo  Add the actual readings.
     *
     * @return array
     */
    private function readingsToImport()
    {
        return [
            // [
            //     'systolic' => 123,
            //     'diastolic' => 81,
            //     'pulse' => 75,
            //     'reading_date' => Carbon::createFromDate(2016, 6, 12),
            //     'created_at' => Carbon::now(),
            //     'updated_at' => Carbon::now(),
            // ],
            [
                'systolic' => 145,
                'diastolic' => 74,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 13), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 119,
                'diastolic' => 64,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 25), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 138,
                'diastolic' => 71,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 26), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 133,
                'diastolic' => 72,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 27), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 131,
                'diastolic' => 72,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 29), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 141,
                'diastolic' => 84,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 30), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 135,
                'diastolic' => 64,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 1), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 129,
                'diastolic' => 74,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 2), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 127,
                'diastolic' => 71,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 6), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 128,
                'diastolic' => 69,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 7), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 122,
                'diastolic' => 67,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 8), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 115,
                'diastolic' => 63,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 9), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 120,
                'diastolic' => 71,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 11), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 129,
                'diastolic' => 65,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 13), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 126,
                'diastolic' => 70,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 14), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 120,
                'diastolic' => 71,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 15), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 122,
                'diastolic' => 66,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 17), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 122,
                'diastolic' => 74,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 18), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 121,
                'diastolic' => 70,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 19), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 124,
                'diastolic' => 63,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 20), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 125,
                'diastolic' => 68,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 21), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
        ];
    }
}

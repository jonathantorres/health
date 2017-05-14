<?php

namespace Tests\Browser;

use App\User;
use Illuminate\Foundation\Testing\DatabaseMigrations;
use Laravel\Dusk\Browser;
use Tests\DuskTestCase;

class LoginTest extends DuskTestCase
{
    public function setUp()
    {
        parent::setUp();

        $usersToDelete = User::where('email', '!=', 'jonathantorres41@gmail.com')->get();

        foreach ($usersToDelete as $userToDelete) {
            $userToDelete->delete();
        }
    }

    /** @test */
    public function existing_user_should_be_able_to_login()
    {
        $user = factory(User::class)->create();
        $this->browse(function (Browser $browser) use ($user) {
            $browser->visit('/login')
                    ->assertSee('Login')
                    ->type('email', $user->email)
                    ->type('password', 'secret')
                    ->press('Login')
                    ->assertPathIs('/')
                    ->assertSee('Dashboard');
        });
    }

    /** @test */
    public function existing_user_should_be_able_to_logout()
    {
        $user = factory(User::class)->create();
        $this->browse(function (Browser $browser) use ($user) {
            $browser->loginAs($user)
                    ->visit('/')
                    ->assertSee('Dashboard')
                    ->clickLink('Logout')
                    ->assertPathIs('/login')
                    ->assertSee('Login');
        });
    }

    public function unexisting_user_should_not_be_able_to_login()
    {
        $this->browse(function (Browser $browser) use ($user) {
            $browser->visit('/login')
                    ->assertSee('Login')
                    ->type('email', 'someone@email.com')
                    ->type('password', 'mypass')
                    ->type('password_confirmation', 'mypass')
                    ->press('Login')
                    ->assertPathIs('/login')
                    ->assertSee('These credentials do not match our records.');
        });
    }
}

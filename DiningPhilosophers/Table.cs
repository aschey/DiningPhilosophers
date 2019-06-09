using System;
using System.Linq;
using System.Collections.Generic;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace DiningPhilosophers
{
    class Table
    {
        List<Philosopher> Philosophers { get; set; }
        List<Fork> Forks { get; set; }

        Waiter Waiter { get; set; }

        public Table()
        {
            var names = new[]
            {
                "Aristotle",
                "Socrates",
                "Confucius",
                "Newton",
                "Locke",
                "Kant",
                "Marx",
                "Nietzsche",
                "Darwin",
                "Descartes",
                "Machiavelli",
                "Hobbes",
                "Chomsky"
            };
            Forks = Enumerable.Range(0, names.Length).Select((i) => new Fork()).ToList();
            
            Waiter = new Waiter();

            Philosophers = new List<Philosopher>();

            for (var i = 0; i < names.Length; i++)
            {
                var nextFork = i + 1;
                if (nextFork == names.Length)
                {
                    nextFork = 0;
                }
                var philosopher = new Philosopher
                {
                   
                    Name = names[i],
                    LeftFork = Forks[i],
                    RightFork = Forks[nextFork],
                    Waiter = Waiter,
                    ThinkTime = 0,
                    EatTime = 1000,
                    ThinkVariance = 0,
                    EatVariance = 0
                };
                Philosophers.Add(philosopher);
            }

            for (var i = 0; i < names.Length; i++)
            {
                var left = i - 1;
                var right = i + 1;
                if (left == -1)
                {
                    left = names.Length - 1;
                }
                if (right == names.Length)
                {
                    right = 0;
                }
                Philosophers[i].LeftPhilosopher = Philosophers[left];
                Philosophers[i].RightPhilosopher = Philosophers[right];
            }

            EventManager.Subscribe("Eating", (name) => Console.WriteLine(Forks.Count(f => f.InUse)), false);
        }

        public void Run()
        {
            Parallel.ForEach(Philosophers, p => p.Run());
            Console.ReadLine();
        }
    }
}
